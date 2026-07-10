package product

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestCreateProductReqValidatesAuctionEnd(t *testing.T) {
	baseReq := CreateProductReq{
		SellerID:    uuid.New(),
		ProductName: "Sample Product",
		Description: "This is a sample product description.",
		Baseprice:   99.99,
	}

	t.Run("rejects auction shorter than two hours", func(t *testing.T) {
		req := baseReq
		req.AuctionEnd = time.Now().Add(minAuctionDuration - time.Minute)

		problems := req.Valid(context.Background())
		if problems["auction_end"] != "must be at least two hours duration" {
			t.Fatalf("expected auction_end validation error, got %v", problems)
		}
	})

	t.Run("accepts auction longer than two hours", func(t *testing.T) {
		req := baseReq
		req.AuctionEnd = time.Now().Add(minAuctionDuration + time.Minute)

		problems := req.Valid(context.Background())
		if _, ok := problems["auction_end"]; ok {
			t.Fatalf("did not expect auction_end validation error, got %v", problems)
		}
	})
}
