package policyhandler

import (
	"fmt"

	"github.com/armosec/kubescape/cautils/getter"
	"github.com/armosec/kubescape/cautils/opapolicy"
)

func (policyHandler *PolicyHandler) GetPoliciesFromBackend(notification *opapolicy.PolicyNotification, getPolicies getter.IPolicyGetter) ([]opapolicy.Framework, error) {
	var errs error
	// d := getter.NewArmoAPI()
	frameworks := []opapolicy.Framework{}
	// Get - cacli opa get
	for _, rule := range notification.Rules {
		switch rule.Kind {
		case opapolicy.KindFramework:
			// backend
			receivedFramework, err := getPolicies.GetFramework(rule.Name)
			if err != nil {
				errs = err
			}
			if receivedFramework != nil {
				frameworks = append(frameworks, *receivedFramework)
			}

		default:
			err := fmt.Errorf("missing rule kind, expected: %s", opapolicy.KindFramework)
			errs = fmt.Errorf("%s", err.Error())
		}
	}
	return frameworks, errs
}
