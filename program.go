package vm

import (
	"unsafe"
)

type Programmer []byte

func (p Programmer) Stop() Programmer {
	return p.Append(Stop)
}

func (p Programmer) Append(op Op, args ...uint64) Programmer {
	b := make([]byte, 2+8*len(args))
	*(*Op)(unsafe.Pointer(&b[0])) = op
	for i, a := range args {
		SetU(a, &b[2+i*8])
	}
	return append(p, b...)
}

func (p Programmer) SetU(r, v uint64) Programmer {
	return p.Append(Set, r, v)
}

func (p Programmer) SetF(r uint64, v float64) Programmer {
	uv := *(*uint64)(unsafe.Pointer(&v))
	return p.Append(Set, r, uv)
}

func (p Programmer) Copy(r1, r2 uint64) Programmer {
	return p.Append(Copy, r1, r2)
}

func (p Programmer) IAdd(r1, r2 uint64) Programmer {
	return p.Append(IAdd, r1, r2)
}

func (p Programmer) FAdd(r1, r2 uint64) Programmer {
	return p.Append(FAdd, r1, r2)
}

func (p Programmer) ISub(r1, r2 uint64) Programmer {
	return p.Append(ISub, r1, r2)
}

func (p Programmer) FSub(r1, r2 uint64) Programmer {
	return p.Append(FSub, r1, r2)
}

func (p Programmer) Alloc(r uint64) Programmer {
	return p.Append(Alloc, r)
}

func (p Programmer) Read(r1, r2, r3 uint64) Programmer {
	return p.Append(Read, r1, r2, r3)
}

func (p Programmer) Write(r1, r2, r3 uint64) Programmer {
	return p.Append(Write, r1, r2, r3)
}

func (p Programmer) Jump(r1, r2, r3 uint64) Programmer {
	return p.Append(Jump, r1, r2, r3)
}

func (p Programmer) Position(r1, r2 uint64) Programmer {
	return p.Append(Position, r1, r2)
}
