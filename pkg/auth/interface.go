package token

import "context"

// Token -
type Token interface {
	// Create - 建立權杖
	Create(ctx context.Context, value interface{}, opts ...Options) (string, error)
}
