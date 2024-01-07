package buffrd

import "context"

type Output interface {
	Write(ctx context.Context, message *message) error
}

type MysqlOutput struct {
}

func (output *MysqlOutput) Write(ctx context.Context, message *message) error {
	// TODO
	return nil
}
