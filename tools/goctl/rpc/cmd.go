package rpc

import (
	"github.com/spf13/cobra"

	"github.com/zeromicro/go-zero/tools/goctl/rpc/cli"
)

var (
	// Cmd describes a rpc command.
	Cmd = &cobra.Command{
		Use:   "rpc",
		Short: "Generate rpc code",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cli.RPCTemplate(true)
		},
	}

	newCmd = &cobra.Command{
		Use:   "new",
		Short: "Generate rpc demo service",
		Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		RunE:  cli.RPCNew,
	}

	templateCmd = &cobra.Command{
		Use:   "template",
		Short: "Generate proto template",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cli.RPCTemplate(false)
		},
	}

	protocCmd = &cobra.Command{
		Use:     "protoc",
		Short:   "Generate grpc code",
		Example: "goctl rpc protoc xx.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=.",
		Args:    cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		RunE:    cli.ZRPC,
	}

	entCmd = &cobra.Command{
		Use:   "ent",
		Short: "Generate Ent CRUD template",
		RunE:  cli.EntCRUDLogic,
	}
)

func init() {
	Cmd.Flags().StringVar(&cli.VarStringOutput, "o", "", "Output a sample proto file")
	Cmd.Flags().StringVar(&cli.VarStringHome, "home", "", "The goctl home path of "+
		"the template, --home and --remote cannot be set at the same time, if they are, --remote has"+
		" higher priority")
	Cmd.Flags().StringVar(&cli.VarStringRemote, "remote", "", "The remote git repo"+
		" of the template, --home and --remote cannot be set at the same time, if they are, --remote"+
		" has higher priority\n\tThe git repo directory must be consistent with the "+
		"https://github.com/zeromicro/go-zero-template directory structure")
	Cmd.Flags().StringVar(&cli.VarStringBranch, "branch", "", "The branch of the "+
		"remote repo, it does work with --remote")

	newCmd.Flags().StringSliceVar(&cli.VarStringSliceGoOpt, "go_opt", nil, "")
	newCmd.Flags().StringSliceVar(&cli.VarStringSliceGoGRPCOpt, "go-grpc_opt", nil, "")
	newCmd.Flags().StringVar(&cli.VarStringStyle, "style", "gozero", "The file "+
		"naming format, see [https://github.com/zeromicro/go-zero/tree/master/tools/goctl/config/readme.md]")
	newCmd.Flags().BoolVar(&cli.VarBoolIdea, "idea", false, "Whether the command "+
		"execution environment is from idea plugin.")
	newCmd.Flags().StringVar(&cli.VarStringHome, "home", "", "The goctl home path "+
		"of the template, --home and --remote cannot be set at the same time, if they are, --remote "+
		"has higher priority")
	newCmd.Flags().StringVar(&cli.VarStringRemote, "remote", "", "The remote git "+
		"repo of the template, --home and --remote cannot be set at the same time, if they are, "+
		"--remote has higher priority\n\tThe git repo directory must be consistent with the "+
		"https://github.com/zeromicro/go-zero-template directory structure")
	newCmd.Flags().StringVar(&cli.VarStringBranch, "branch", "",
		"The branch of the remote repo, it does work with --remote")
	newCmd.Flags().BoolVarP(&cli.VarBoolVerbose, "verbose", "v", false, "Enable log output")
	newCmd.Flags().BoolVar(&cli.VarBoolEnt, "ent", true, "Whether use Ent in project")
	newCmd.Flags().MarkHidden("go_opt")
	newCmd.Flags().MarkHidden("go-grpc_opt")
	newCmd.Flags().StringVar(&cli.VarStringModuleName, "module_name", "",
		"The module name in go.mod. e.g. github.com/suyuan32/simple-admin-core")
	newCmd.Flags().StringVar(&cli.VarStringGoZeroVersion, "go_zero_version", "",
		"The go zero version used for replacement. e.g. v1.4.3")
	newCmd.Flags().StringVar(&cli.VarStringToolVersion, "tool_version", "",
		"The simple admin tool version version used for migration. e.g. v0.1.7")
	newCmd.Flags().IntVar(&cli.VarIntServicePort, "port", 9110, "The service port exposed")
	newCmd.Flags().BoolVar(&cli.VarBoolGitlab, "gitlab", false, "Whether to use gitlab-ci")
	newCmd.Flags().BoolVar(&cli.VarBoolDesc, "desc", false, "Whether to create desc folder for splitting proto files")

	protocCmd.Flags().BoolVarP(&cli.VarBoolMultiple, "multiple", "m", false,
		"Generated in multiple rpc service mode")
	protocCmd.Flags().StringSliceVar(&cli.VarStringSliceGoOut, "go_out", nil, "")
	protocCmd.Flags().StringSliceVar(&cli.VarStringSliceGoGRPCOut, "go-grpc_out", nil, "")
	protocCmd.Flags().StringSliceVar(&cli.VarStringSliceGoOpt, "go_opt", nil, "")
	protocCmd.Flags().StringSliceVar(&cli.VarStringSliceGoGRPCOpt, "go-grpc_opt", nil, "")
	protocCmd.Flags().StringSliceVar(&cli.VarStringSlicePlugin, "plugin", nil, "")
	protocCmd.Flags().StringSliceVarP(&cli.VarStringSliceProtoPath, "proto_path", "I", nil, "")
	protocCmd.Flags().StringVar(&cli.VarStringZRPCOut, "zrpc_out", "", "The zrpc output directory")
	protocCmd.Flags().StringVar(&cli.VarStringStyle, "style", "go_zero", "The file "+
		"naming format, see [https://github.com/zeromicro/go-zero/tree/master/tools/goctl/config/readme.md]")
	protocCmd.Flags().StringVar(&cli.VarStringHome, "home", "", "The goctl home "+
		"path of the template, --home and --remote cannot be set at the same time, if they are, "+
		"--remote has higher priority")
	protocCmd.Flags().StringVar(&cli.VarStringRemote, "remote", "", "The remote "+
		"git repo of the template, --home and --remote cannot be set at the same time, if they are, "+
		"--remote has higher priority\n\tThe git repo directory must be consistent with the "+
		"https://github.com/zeromicro/go-zero-template directory structure")
	protocCmd.Flags().StringVar(&cli.VarStringBranch, "branch", "",
		"The branch of the remote repo, it does work with --remote")
	protocCmd.Flags().BoolVarP(&cli.VarBoolVerbose, "verbose", "v", false, "Enable log output")
	protocCmd.Flags().MarkHidden("go_out")
	protocCmd.Flags().MarkHidden("go-grpc_out")
	protocCmd.Flags().MarkHidden("go_opt")
	protocCmd.Flags().MarkHidden("go-grpc_opt")
	protocCmd.Flags().MarkHidden("plugin")
	protocCmd.Flags().MarkHidden("proto_path")

	templateCmd.Flags().StringVar(&cli.VarStringOutput, "o", "", "Output a sample proto file")
	templateCmd.Flags().StringVar(&cli.VarStringHome, "home", "", "The goctl home"+
		" path of the template, --home and --remote cannot be set at the same time, if they are, "+
		"--remote has higher priority")
	templateCmd.Flags().StringVar(&cli.VarStringRemote, "remote", "", "The remote "+
		"git repo of the template, --home and --remote cannot be set at the same time, if they are, "+
		"--remote has higher priority\n\tThe git repo directory must be consistent with the "+
		"https://github.com/zeromicro/go-zero-template directory structure")
	templateCmd.Flags().StringVar(&cli.VarStringBranch, "branch", "", "The branch"+
		" of the remote repo, it does work with --remote")

	entCmd.Flags().StringVar(&cli.VarStringSchema, "schema", "", "The schema path of the Ent")
	entCmd.Flags().StringVar(&cli.VarStringOutput, "o", "", "The output path")
	entCmd.Flags().StringVar(&cli.VarStringServiceName, "service_name", "", "The service name")
	entCmd.Flags().StringVar(&cli.VarStringProjectName, "project_name", "", "The project name")
	entCmd.Flags().BoolVar(&cli.VarBoolMultiple, "multiple", false, "Generated in multiple rpc service mode")
	entCmd.Flags().StringVar(&cli.VarStringStyle, "style", "go_zero", "The file name format style")
	entCmd.Flags().StringVar(&cli.VarStringModelName, "model", "", "The model name for generating e.g. user, "+
		"if it is empty, generate codes for all models in schema directory")
	entCmd.Flags().IntVar(&cli.VarIntSearchKeyNum, "search_key_num", 3, "The max number of search keys ")
	entCmd.Flags().StringVar(&cli.VarStringGroupName, "group", "", "The group name for logic. e.g. user")
	entCmd.Flags().StringVar(&cli.VarStringProtoPath, "proto_out", "", "The output proto file path")

	Cmd.AddCommand(newCmd)
	Cmd.AddCommand(protocCmd)
	Cmd.AddCommand(templateCmd)
	Cmd.AddCommand(entCmd)
}
