package repository

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/Masterminds/squirrel"
)

// Sqlizer ...
type Sqlizer interface {
	ToSql() (sql string, args []interface{}, err error)
}

type toSQLFn func() (sqlStr string, args []interface{}, err error)

func (fn toSQLFn) ToSql() (sqlStr string, args []interface{}, err error) { return fn() }

//TODO подумать над необходимостью
//// Exec ...
//func Exec (
//	ctx context.Context,
//	e interface{
//		Exec(ctx context.Context, sqlStr string, args ...interface{}) (sql.Result, error)
//},
//query string,
//args ...interface{},
//) (sql.Result, error) {
//	cmd, err := e.Exec(ctx, query, args...)
//	return cmd, err
//}

// Select упрощенная версия без mapToStruct
func Select[TSlice ~[]*T, T any](
	ctx context.Context,
	e interface {
		Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	},
	dest *TSlice,
	query string,
	args ...interface{},
) error {
	if dest == nil {
		return fmt.Errorf("dest cannot be nil")
	}

	rows, err := e.Query(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	return scanRowsToSlice(rows, dest)
}

func Selectx[TSlice ~[]*T, T any](
	ctx context.Context,
	e interface {
		Query(ctx context.Context, sql string, args ...interface{}) (*sql.Rows, error)
	},
	dest *TSlice,
	sqlizer Sqlizer,
) error {
	sqlizer = ReplacePlaceholders(sqlizer)
	stmt, args, err := sqlizer.ToSql()
	if err != nil {
		return err
	}

	return Select[TSlice](ctx, e, dest, stmt, args...)
}

func ReplacePlaceholders(sqlizer Sqlizer) Sqlizer {
	var (
		sqlStr string
		args   []interface{}
		err    error
	)

	fn := toSQLFn(func() (string, []interface{}, error) { return sqlStr, args, err })

	sqlStr, args, err = sqlizer.ToSql()
	if err != nil {
		return fn
	}

	sqlStr, err = squirrel.Dollar.ReplacePlaceholders(sqlStr)

	return fn
}

// scanRowsToSlice сканирует строки в срез структур
func scanRowsToSlice[TSlice ~[]*T, T any](rows *sql.Rows, dest *TSlice) error {
	// Получаем колонки
	columns, err := rows.Columns()
	if err != nil {
		return fmt.Errorf("get columns failed: %w", err)
	}

	// Кэш для информации о полях
	var fieldCache sync.Map

	// Создаем буфер для сканирования
	scanArgs := make([]interface{}, len(columns))
	for i := range scanArgs {
		scanArgs[i] = new(interface{})
	}

	var results []*T

	for rows.Next() {
		if err := rows.Scan(scanArgs...); err != nil {
			return fmt.Errorf("scan failed: %w", err)
		}

		// Создаем новую структуру
		item := new(T)

		// Получаем или создаем маппинг полей для этого типа
		var t T
		structType := reflect.TypeOf(t)

		// Пытаемся получить из кэша
		cacheKey := structType.String()
		var columnMapping []int
		if cached, ok := fieldCache.Load(cacheKey); ok {
			columnMapping = cached.([]int)
		} else {
			// Создаем маппинг
			columnMapping = make([]int, len(columns))
			for i, col := range columns {
				columnMapping[i] = -1 // По умолчанию не найдено

				// Ищем поле в структуре
				colLower := strings.ToLower(col)
				for j := 0; j < structType.NumField(); j++ {
					field := structType.Field(j)
					if !field.IsExported() {
						continue
					}

					// Проверяем тег db
					fieldName := field.Name
					if dbTag := field.Tag.Get("db"); dbTag != "" && dbTag != "-" {
						// Берем первую часть тега
						if commaIdx := strings.Index(dbTag, ","); commaIdx != -1 {
							fieldName = dbTag[:commaIdx]
						} else {
							fieldName = dbTag
						}
					}

					if strings.ToLower(fieldName) == colLower {
						columnMapping[i] = j
						break
					}
				}
			}
			fieldCache.Store(cacheKey, columnMapping)
		}

		// Заполняем поля структуры
		v := reflect.ValueOf(item).Elem()
		for i, fieldIndex := range columnMapping {
			if fieldIndex == -1 {
				continue // Поле не найдено
			}

			valPtr := scanArgs[i].(*interface{})
			if valPtr == nil {
				continue
			}

			val := reflect.ValueOf(*valPtr)
			if val.IsValid() {
				field := v.Field(fieldIndex)
				if field.CanSet() {
					setFieldValue(field, val)
				}
			}
		}

		results = append(results, item)
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("rows iteration failed: %w", err)
	}

	*dest = results
	return nil
}

// setFieldValue устанавливает значение поля с простой конвертацией
func setFieldValue(field reflect.Value, value reflect.Value) {
	if !field.CanSet() || !value.IsValid() {
		return
	}

	// Если типы совпадают
	if value.Type().AssignableTo(field.Type()) {
		field.Set(value)
		return
	}

	// Простые конвертации
	switch field.Kind() {
	case reflect.String:
		field.SetString(fmt.Sprint(value.Interface()))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch value.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			field.SetInt(value.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			field.SetInt(int64(value.Uint()))
		case reflect.Float32, reflect.Float64:
			field.SetInt(int64(value.Float()))
		case reflect.String:
			var intVal int64
			fmt.Sscanf(value.String(), "%d", &intVal)
			field.SetInt(intVal)
		}
	case reflect.Bool:
		switch value.Kind() {
		case reflect.Bool:
			field.SetBool(value.Bool())
		case reflect.String:
			str := strings.ToLower(value.String())
			field.SetBool(str == "true" || str == "1" || str == "t")
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			field.SetBool(value.Int() != 0)
		}
	}
}
