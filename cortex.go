package cortex

import (
	//"encoding/json"
	//"errors"
	//"fmt"
	"log"
	"math/rand"
	//"net/http"
	//"net/url"
	//"strings"
	"time"

	"github.com/itsabot/abot/shared/datatypes"
	//"github.com/itsabot/abot/shared/language"
	"github.com/itsabot/abot/shared/nlp"
	"github.com/itsabot/abot/shared/plugin"
)

type weatherJSON struct {
	Description []string
	Temp        float64
	Humidity    int
}

var p *dt.Plugin

func init() {
	rand.Seed(time.Now().UnixNano())
	trigger := &nlp.StructuredInput{
		Commands: []string{"who"},
		Objects: []string{"you"},
	}
	fns := &dt.PluginFns{Run: Run, FollowUp: FollowUp}
	var err error
	p, err = plugin.New("github.com/crazytweek/plugin_cortex", trigger, fns)
	if err != nil {
		log.Fatal(err)
	}
	p.Vocab = dt.NewVocab(
		dt.VocabHandler{
			Fn: kwIAm,
			Trigger: &nlp.StructuredInput{
				Commands: []string{"who"},
				Objects: []string{"you"},
			},
		},
	)
}

func Run(in *dt.Msg) (string, error) {
	return FollowUp(in)
}

func FollowUp(in *dt.Msg) (string, error) {
	return p.Vocab.HandleKeywords(in), nil
}

func kwIAm(in *dt.Msg) (resp string) {
	return "I am Jet, your personal assistant!"
}

func buildStateMachine(in *dt.Msg) *dt.StateMachine {
	sm := dt.NewStateMachine(p)
	sm.SetStates([]dt.State{})
	sm.LoadState(in)
	return sm
}

func er(err error) string {
	p.Log.Debug(err)
	return "Something went wrong, but I'll try to get that fixed right away."
}
