package postgresrep

func (r *repository) CloseConnection() {
	r.pool.Close()
}
