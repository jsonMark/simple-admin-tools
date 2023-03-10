package {{.modelNameLowerCase}}

import (
	"context"
	"net/http"

	"{{.projectPackage}}/internal/svc"
	"{{.projectPackage}}/internal/types"
	"{{.rpcPackage}}"

	"github.com/zeromicro/go-zero/core/logx"
)

type Update{{.modelName}}Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewUpdate{{.modelName}}Logic(r *http.Request, svcCtx *svc.ServiceContext) *Update{{.modelName}}Logic {
	return &Update{{.modelName}}Logic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *Update{{.modelName}}Logic) Update{{.modelName}}(req *types.{{.modelName}}Info) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.{{.rpcName}}Rpc.Update{{.modelName}}(l.ctx,
		&{{.rpcPbPackageName}}.{{.modelName}}Info{
			Id:          req.Id,{{.setLogic}}
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.lang, data.Msg)}, nil
}
