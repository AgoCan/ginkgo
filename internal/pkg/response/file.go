package response

const (
	CreateDirectorErr = 54001
	CreateFileErr     = 54002
)

func fileRes(msg map[int]string) map[int]string {
	msg[CreateDirectorErr] = "创建目录失败"
	msg[CreateFileErr] = "创建文件失败"
	return msg
}
