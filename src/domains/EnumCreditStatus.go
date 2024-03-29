package domains

type CreditStatus string

const (
	PRE_PROCESSAMENTO                  CreditStatus = "PRE_PROCESSAMENTO"
	PRE_APROVADO                       CreditStatus = "PRE_APROVADO"
	PENDENTE                           CreditStatus = "PENDENTE"
	AGUARDANDO_DOCUMENTOS              CreditStatus = "AGUARDANDO_DOCUMENTOS"
	EM_ANALISE_MANUAL                  CreditStatus = "EM_ANALISE_MANUAL"
	EM_ANALISE_OPERACIONAL             CreditStatus = "EM_ANALISE_OPERACIONAL"
	EM_ANALISE_DOCUMENTAL              CreditStatus = "EM_ANALISE_DOCUMENTAL"
	CONTRATADO                         CreditStatus = "CONTRATADO"
	LIBERADO                           CreditStatus = "LIBERADO"
	REPROVADO                          CreditStatus = "REPROVADO"
	EXPIRADO                           CreditStatus = "EXPIRADO"
)
