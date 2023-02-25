package delivery

import (
	"bytes"
	"encoding/binary"
	"net"

	"github.com/kuzkuss/iproto_server/internal/usecase"
	"github.com/kuzkuss/iproto_server/models"
	"github.com/vmihailenco/msgpack/v5"
)

type DeliveryI interface {
	GetRequest(conn *net.TCPConn) (*models.Request, error)
	HandleRequest(request *models.Request) (interface{}, error)
}

type delivery struct {
	useCase usecase.UseCaseI
}

func New(uc usecase.UseCaseI) DeliveryI {
	return &delivery{
		useCase: uc,
	}
}

func (del *delivery) GetRequest(conn *net.TCPConn) (*models.Request, error) {
	headerBytes := make([]byte, 12)
	_, err := conn.Read(headerBytes)
	if err != nil {
		return nil, err
	}

	req := models.Request{}
	headerBuf := bytes.NewBuffer(headerBytes)

	if err := binary.Read(headerBuf, binary.LittleEndian, &req.Header.FuncId); err != nil {
		return nil, err
	}

	if err := binary.Read(headerBuf, binary.LittleEndian, &req.Header.BodyLength); err != nil {
		return nil, err
	}

	if err := binary.Read(headerBuf, binary.LittleEndian, &req.Header.RequestId); err != nil {
		return nil, err
	}

	req.Body = make([]byte, req.Header.BodyLength)

	_, err = conn.Read(req.Body)
	if err != nil {
		return nil, err
	}

	return &req, nil
}

func (del *delivery) HandleRequest(request *models.Request) (interface{}, error) {
	switch request.Header.FuncId {
	case models.ADM_STORAGE_SWITCH_READONLY:
		err := del.useCase.SwitchState(models.READ_ONLY)
		if err != nil {
			return err.Error(), err
		}
		return nil, nil

	case models.ADM_STORAGE_SWITCH_READWRITE:
		err := del.useCase.SwitchState(models.READ_WRITE)
		if err != nil {
			return err.Error(), err
		}
		return nil, nil
	
	case models.ADM_STORAGE_SWITCH_MAINTENANCE:
		err := del.useCase.SwitchState(models.MAINTENANCE)
		if err != nil {
			return err.Error(), err
		}
		return nil, nil
		
	case models.STORAGE_REPLACE:
		req := models.RequestSaveString{}
		err := msgpack.Unmarshal(request.Body, &req)
		if err != nil {
			return err.Error(), err
		}
		err = del.useCase.SaveString(req.Idx, req.Str)
		if err != nil {
			return err.Error(), err
		}
		return nil, nil
		
	case models.STORAGE_READ:
		req := models.RequestGetString{}
		err := msgpack.Unmarshal(request.Body, &req)
		if err != nil {
			return err.Error(), err
		}
		str, err := del.useCase.GetString(req.Idx)
		if err != nil {
			return err.Error(), err
		}
		return str, nil
	
	default:
		return nil, models.ErrFuncId
	}
}

