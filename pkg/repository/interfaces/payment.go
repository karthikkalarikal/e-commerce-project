package interfaces

type PaymentRepository interface {
	AddRazorPayDetails(int, string) error
}
