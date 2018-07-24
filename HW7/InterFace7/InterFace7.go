package InterFace7

type Phone interface {
	Sensor() string
	Monitor() string
	Button() string
}

func Unlock(p Phone) {
	println("已使用" + p.Sensor() + "解鎖")
	println("Hi~ " + p.Monitor())
}

func Lock(p Phone) {
	println("已用" + p.Button() + "上鎖")
}

type PhoneZ interface {
	Attack() string
	Run() string
	Move() string
}

func Fire(z PhoneZ) {
	println("已使用" + z.Attack() + "來攻擊鎖")
	println("用 " + z.Run() + "來前進")
}

func Miss(z PhoneZ) {
	println("已用" + z.Move() + "來閃避")
}
