package main

import (
	"context"
	"fmt"
)

type ck string // custom key

func hasKey(ctx context.Context, key ck) bool {
	if v := ctx.Value(key); v != nil {
		return true
	}
	return false
}

func main() {
	idKey := ck("id")
	emailKey := ck("email")
	secretKey := ck("secret")

	// parent context
	ctx := context.Background()

	// child context
	ctx = context.WithValue(ctx, idKey, 1)

	// child context
	ctx = context.WithValue(ctx, emailKey, "vigo@foo.com")

	fmt.Println("idKey", hasKey(ctx, idKey))
	fmt.Println("emailKey", hasKey(ctx, emailKey))
	fmt.Println("secretKey", hasKey(ctx, secretKey))

	if hasKey(ctx, idKey) {
		fmt.Println("value of id", ctx.Value(idKey))
	}
	if hasKey(ctx, emailKey) {
		fmt.Println("value of email", ctx.Value(emailKey))
	}
}
