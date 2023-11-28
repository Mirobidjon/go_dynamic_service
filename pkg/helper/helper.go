package helper

import (
	"database/sql"
	"sort"
	"strconv"
	"strings"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/saidamir98/udevs_pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type queryParams struct {
	Key string
	Val interface{}
}

func ReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var (
		i    int = 1
		args []interface{}
		arr  []queryParams
	)

	for k, v := range params {
		arr = append(arr, queryParams{
			Key: k,
			Val: v,
		})
	}

	sort.Slice(arr, func(i, j int) bool {
		return len(arr[i].Key) > len(arr[j].Key)
	})

	for _, v := range arr {
		if v.Key != "" && strings.Contains(namedQuery, ":"+v.Key) {
			namedQuery = strings.ReplaceAll(namedQuery, ":"+v.Key, "$"+strconv.Itoa(i))
			args = append(args, v.Val)
			i++
		}
	}

	return namedQuery, args
}

func ReplaceSQL(old, searchPattern string) string {
	tmpCount := strings.Count(old, searchPattern)
	for m := 1; m <= tmpCount; m++ {
		old = strings.Replace(old, searchPattern, "$"+strconv.Itoa(m), 1)
	}
	return old
}

func HandleError(log logger.LoggerI, err error, message string, req interface{}, code codes.Code) error {
	if code != 0 && err != nil {
		log.Error(message, logger.Error(err), logger.Any("req", req))
		return status.Error(code, message+err.Error())
	} else if err == sql.ErrNoRows {
		log.Error(message+", Not Found", logger.Error(err), logger.Any("req", req))
		return status.Error(codes.NotFound, "Not Found")
	} else if err != nil {
		log.Error(message, logger.Error(err), logger.Any("req", req))
		return status.Error(codes.Internal, message+err.Error())
	}

	log.Error(message, logger.Error(err), logger.Any("req", req))
	return status.Error(codes.Internal, message)
}

func ToNullString(s *wrappers.StringValue) (res sql.NullString) {
	if s.GetValue() != "" {
		res.String = s.Value
		res.Valid = true
	}
	return res
}

func ToStringValue(s sql.NullString) *wrappers.StringValue {
	if s.Valid {
		return &wrappers.StringValue{Value: s.String}
	}
	return nil
}

func TotalSumma(totalsum *float64, s string) error {
	sum, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	*totalsum += sum
	return nil
}
