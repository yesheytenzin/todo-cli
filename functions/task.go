package functions

type Task struct {
	ID				int				`json:"id"`
	Title			string		`json:"title"`
	Completed	bool			`json:"completed"`
}
