package protocgentoolservice

type ProtocGenToolService interface {
	Generate() error
}

/**
 * @description: CreateProtoGenToolService
 * @param {ProtocGenToolService} ProtocGenToolService
 * @author: Jerry.Yang
 * @date: 2023-12-12 11:36:28
 * @return {*}
 */
func CreateProtoGenToolService(ProtocGenToolService ProtocGenToolService) ProtocGenToolService {
	return ProtocGenToolService
}
