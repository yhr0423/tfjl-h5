package protocols

import (
	"bytes"
	"encoding/binary"
)

type C_Network_Fight_To_Logic_Role_FightBalance struct {
	RoleID            int64
	SRoleFightBalance S_Role_FightBalance
}

func (p *C_Network_Fight_To_Logic_Role_FightBalance) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.RoleID)
	buffer.Write(p.SRoleFightBalance.Encode())
	return buffer.Bytes()
}

func (p *C_Network_Fight_To_Logic_Role_FightBalance) Decode(buffer *bytes.Buffer) error {
	binary.Read(buffer, binary.LittleEndian, &p.RoleID)
	p.SRoleFightBalance.Decode(buffer)
	return nil
}
