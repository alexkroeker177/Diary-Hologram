// This file contains the implementation of Diary.

package app

import (
	"diary/util"
	"fmt"
	"github.com/Bitspark/go-bitnode/bitnode"
	"reflect"
)

// Struct definition for Diary.

// Diary is the main sparkable.
type Diary struct {
	bitnode.System

	// EntryList description: Value of entryList in Diary.
	EntryList []Entry `json:"entryList"`
}

// Diary methods.

// AddEntry description: Add an entry to the diary.
func (s *Diary) AddEntry(entry Entry) error {
	// TODO: Implement method.
	return fmt.Errorf("method addEntry not implemented yet")
}

// RemoveEntry description: Remove an entry from the diary.
func (s *Diary) RemoveEntry(id string) error {
	// TODO: Implement method.
	return fmt.Errorf("method removeEntry not implemented yet")
}

// AddTag description: Method addTag of Diary.
func (s *Diary) AddTag(entry Entry, tag Tag) error {
	// TODO: Implement method.
	return fmt.Errorf("method addTag not implemented yet")
}

// HandleEntryList reacts on changes of the entryList variable.
func (s *Diary) HandleEntryList(entryList []Entry) {
	// TODO: Implement handler.
}

// Lifecycle callbacks.

// lifecycleCreate is called when the container has been created.
func (s *Diary) lifecycleCreate(vals ...bitnode.HubItem) error {
	// TODO: Add startup logic here which is called when the spark is created.
	return nil
}

// lifecycleLoad is called when the container has been started (after lifecycleCreate) or restarted.
func (s *Diary) lifecycleLoad(vals ...bitnode.HubItem) error {
	// TODO: Add startup logic here which is called after the spark has been created.

	s.SetMessage("Diary running...")
	s.SetStatus(bitnode.SystemStatusRunning)

	return nil
}

// DO NOT CHANGE THE FOLLOWING CODE UNLESS YOU KNOW WHAT YOU ARE DOING.

func (s *Diary) Update(values ...string) error {
	sv := reflect.ValueOf(*s)
	st := reflect.TypeOf(*s)
	if len(values) == 0 {
		for i := 0; i < st.NumField(); i++ {
			values = append(values, st.Field(i).Name)
		}
	}
	for _, value := range values {
		ft, ok := st.FieldByName(value)
		if !ok {
			return fmt.Errorf("field '%s' not found in %s", value, st.Name())
		}
		fv := sv.FieldByName(value)
		if !fv.IsValid() {
			return fmt.Errorf("field '%s' not found in %s", value, st.Name())
		}
		val, err := util.InterfaceFromValue(fv.Interface())
		if err != nil {
			return err
		}
		hubName := ft.Tag.Get("json")
		if err := s.GetHub(hubName).Set("", val); err != nil {
			return err
		}
	}
	return nil
}

// Init attaches the methods of the Diary to the respective handlers.
func (s *Diary) Init() error {
	// METHODS

	s.GetHub("addEntry").Handle(bitnode.NewNativeFunction(func(creds bitnode.Credentials, vals ...bitnode.HubItem) ([]bitnode.HubItem, error) {
		entry, err := util.ValueFromInterface[Entry](vals[0])
		if err != nil {
			return nil, err
		}

		err = s.AddEntry(entry)
		if err != nil {
			return nil, err
		}

		return []bitnode.HubItem{}, nil
	}))

	s.GetHub("removeEntry").Handle(bitnode.NewNativeFunction(func(creds bitnode.Credentials, vals ...bitnode.HubItem) ([]bitnode.HubItem, error) {
		id, err := util.ValueFromInterface[string](vals[0])
		if err != nil {
			return nil, err
		}

		err = s.RemoveEntry(id)
		if err != nil {
			return nil, err
		}

		return []bitnode.HubItem{}, nil
	}))

	s.GetHub("addTag").Handle(bitnode.NewNativeFunction(func(creds bitnode.Credentials, vals ...bitnode.HubItem) ([]bitnode.HubItem, error) {
		entry, err := util.ValueFromInterface[Entry](vals[0])
		if err != nil {
			return nil, err
		}
		tag, err := util.ValueFromInterface[Tag](vals[1])
		if err != nil {
			return nil, err
		}

		err = s.AddTag(entry, tag)
		if err != nil {
			return nil, err
		}

		return []bitnode.HubItem{}, nil
	}))

	// VALUES

	s.GetHub("entryList").Subscribe(bitnode.NewNativeSubscription(func(id string, creds bitnode.Credentials, val bitnode.HubItem) {
		entryList, err := util.ValueFromInterface[[]Entry](val)
		if err != nil {
			return
		}
		s.EntryList = entryList
		s.HandleEntryList(entryList)
	}))

	// CHANNELS

	// LIFECYCLE EVENTS

	s.AddCallback(bitnode.LifecycleCreate, bitnode.NewNativeEvent(func(vals ...bitnode.HubItem) error {
		return s.lifecycleCreate(vals...)
	}))

	s.AddCallback(bitnode.LifecycleLoad, bitnode.NewNativeEvent(func(vals ...bitnode.HubItem) error {
		return s.lifecycleLoad(vals...)
	}))

	return nil
}
