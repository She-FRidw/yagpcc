package gp

import (
	"context"
	"strings"
)

func IsCloudberryOrGP7(ctx context.Context) (bool, error) {
	version, err := GetVersion(ctx)
	if err != nil {
		return false, err
	}
	if strings.Contains(version.Version, "loudberry") ||
		strings.Contains(version.Version, "Greenplum Database 7") {
		return true, nil
	}
	return false, nil
}
