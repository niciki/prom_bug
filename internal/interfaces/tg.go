// @tg version=1.0.0
// @tg packageJSON=`github.com/seniorGolang/json`
//
//go:generate tg transport --services . --out ../transport --outSwagger ../../api/swagger.yaml
//go:generate tg client --services . --outPath ../clients/users

package interfaces
