package gp

import (
	"context"
	"fmt"
)

type StatActivityColumnsConfiguration struct {
	HasWaitEvent bool `db:"has_wait_event"`
}

const StatActivityColumnsQ = `
	SELECT
		count(*) FILTER (WHERE attname = 'wait_event') > 0 AS has_wait_event
	FROM pg_attribute
	WHERE attrelid = 'pg_catalog.pg_stat_activity'::regclass
		AND attnum > 0
		AND NOT attisdropped
`

func UsesModernStatActivity(ctx context.Context) (bool, error) {
	if db == nil {
		return false, fmt.Errorf("internal - DB not initialized")
	}
	statActivityColumns := make([]StatActivityColumnsConfiguration, 0)
	err := db.ExecQuery(ctx, StatActivityColumnsQ, &statActivityColumns)
	if err != nil {
		return false, err
	}
	if len(statActivityColumns) == 0 {
		return false, fmt.Errorf("internal - empty pg_stat_activity columns query result")
	}
	return statActivityColumns[0].HasWaitEvent, nil
}
