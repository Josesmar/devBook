package modelos

//Senha representa o formato da requisição de alteração da senha
type Senha struct {
	Nova  string `json:"nova"`
	Atual string `json:"atual"`
}
