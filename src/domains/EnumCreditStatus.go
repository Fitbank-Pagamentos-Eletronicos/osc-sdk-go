package domains

type CreditStatus struct {
	creditStatusId string
}

var (
	PRE_PROCESSAMENTO                  CreditStatus = CreditStatus{"PRE_PROCESSAMENTO"}
	PRE_APROVADO                       CreditStatus = CreditStatus{"PRE_APROVADO"}
	PENDENTE                           CreditStatus = CreditStatus{"PENDENTE"}
	AGUARDANDO_DOCUMENTOS              CreditStatus = CreditStatus{"AGUARDANDO_DOCUMENTOS"}
	EM_ANALISE_MANUAL                  CreditStatus = CreditStatus{"EM_ANALISE_MANUAL"}
	EM_ANALISE_OPERACIONAL             CreditStatus = CreditStatus{"EM_ANALISE_OPERACIONAL"}
	EM_ANALISE_DOCUMENTAL              CreditStatus = CreditStatus{"EM_ANALISE_DOCUMENTAL"}
	CONTRATADO                         CreditStatus = CreditStatus{"CONTRATADO"}
	LIBERADO                           CreditStatus = CreditStatus{"LIBERADO"}
	REPROVADO                          CreditStatus = CreditStatus{"REPROVADO"}
	EXPIRADO                           CreditStatus = CreditStatus{"EXPIRADO"}
)
