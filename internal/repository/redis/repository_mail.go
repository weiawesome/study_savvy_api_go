package redis

import (
	"context"
	"errors"
)

func (r *Repository) ValidateInVerification(user string, code string) error {
	ctx := context.Background()
	value, err := r.client.Get(ctx, user).Result()
	if err != nil {
		return err
	}

	if value == code {
		return nil
	} else {
		return errors.New("error code")
	}

}
