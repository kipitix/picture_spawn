package arguments

import "strings"

func ParseEtcdEndpoints(srcEndpoints string) []string {
	return strings.Split(srcEndpoints, ",")
}
