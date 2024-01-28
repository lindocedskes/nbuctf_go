package system

const initOrderMenuAuthority = initOrderMenu + initOrderAuthority

type initMenuAuthority struct{}

// auto run
func init() {
	//system.RegisterInit(initOrderMenuAuthority, &initMenuAuthority{})
}

//todo: 未完成
