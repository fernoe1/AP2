package service

import (
	"context"

	"github.com/fernoe1/AP2/assignment-1/payment/internal/domain"
	svc "github.com/fernoe1/protogen/ap2-assign2/service/payment"
	paymentpb "github.com/fernoe1/protogen/ap2-assign2/shares/payment"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PaymentService struct {
	svc.UnimplementedPaymentServiceServer
	UC PaymentUsecase
}

func (ps *PaymentService) GetPaymentByOrderID(ctx context.Context, request *svc.GetPaymentByOrderIDRequest) (*svc.GetPaymentByOrderIDResponse, error) {
	if request.OrderId == "" {
		return nil, status.Error(codes.InvalidArgument, "order_id is required")
	}

	payment, err := ps.UC.GetPaymentFromOrderId(ctx, request.OrderId)
	if err != nil {
		return nil, err
	}

	pbPayment := make([]*paymentpb.Payment, 0, len(payment))
	for _, p := range payment {
		pbPayment = append(pbPayment, &paymentpb.Payment{
			Id:            uint64(p.ID),
			OrderId:       p.OrderID,
			TransactionId: p.TransactionID,
			Amount:        p.Amount,
			Status:        p.Status,
		})
	}

	return &svc.GetPaymentByOrderIDResponse{Payment: pbPayment}, nil
}

func (ps *PaymentService) CreatePayment(ctx context.Context, request *svc.CreatePaymentRequest) (*svc.CreatePaymentResponse, error) {
	if request.OrderId == "" {
		return nil, status.Error(codes.InvalidArgument, "order_id is required")
	}

	if request.Amount < 1 {
		return nil, status.Error(codes.InvalidArgument, "amount > 0")
	}

	payment := &domain.Payment{OrderID: request.OrderId, Amount: request.Amount}
	if err := ps.UC.CreatePayment(ctx, payment); err != nil {
		return nil, err
	}

	pbpayment := &paymentpb.Payment{
		Id:            uint64(payment.ID),
		OrderId:       payment.OrderID,
		TransactionId: payment.TransactionID,
		Amount:        payment.Amount,
		Status:        payment.Status,
	}

	return &svc.CreatePaymentResponse{Payment: pbpayment}, nil
}
