package asset

import (
	"github.com/odpf/compass/core/user"
	"github.com/peterbourgon/mergemap"
)

// patch appends asset with data from map. It mutates the asset itself.
func patchAsset(a *Asset, patchData map[string]interface{}) {
	a.URN = patchString("urn", patchData, a.URN)
	a.URL = patchString("url", patchData, a.URL)
	a.Type = Type(patchString("type", patchData, a.Type.String()))
	a.Service = patchString("service", patchData, a.Service)
	a.Name = patchString("name", patchData, a.Name)
	a.Description = patchString("description", patchData, a.Description)

	labels, exists := patchData["labels"]
	if exists {
		a.Labels = buildLabels(labels)
	}
	owners, exists := patchData["owners"]
	if exists {
		a.Owners = buildOwners(owners)
	}
	data, exists := patchData["data"]
	if exists {
		patchAssetData(a, data)
	}
	attributes, exists := patchData["attributes"]
	if exists {
		patchAssetAttributes(a, attributes)
	}
}

// buildLabels builds labels from interface{}
func buildLabels(data interface{}) (labels map[string]string) {
	switch d := data.(type) {
	case map[string]interface{}:
		labels = map[string]string{}
		for key, value := range d {
			stringVal, ok := value.(string)
			if !ok {
				continue
			}
			labels[key] = stringVal
		}
	case map[string]string:
		labels = d
	default:
		labels = nil
	}

	return
}

// buildOwners builds owners from interface{}
func buildOwners(data interface{}) (owners []user.User) {
	buildOwner := func(data map[string]interface{}) user.User {
		return user.User{
			ID:       getString("id", data),
			Email:    getString("email", data),
			Provider: getString("provider", data),
		}
	}

	switch d := data.(type) {
	case []interface{}:
		owners = []user.User{}
		for _, value := range d {
			mapValue, ok := value.(map[string]interface{})
			if !ok {
				continue
			}
			owners = append(owners, buildOwner(mapValue))
		}
	case []map[string]interface{}:
		owners = []user.User{}
		for _, value := range d {
			owners = append(owners, buildOwner(value))
		}
	case []user.User:
		owners = d
	default:
		owners = nil
	}

	return
}

// patchAssetData patches asset's data using map
func patchAssetData(a *Asset, data interface{}) {
	if data == nil {
		return
	}
	dataMap, ok := data.(map[string]interface{})
	if !ok {
		return
	}

	if a.Data == nil {
		a.Data = dataMap
		return
	}

	a.Data = mergemap.Merge(a.Data, dataMap)
}

// patchAssetAttributes patches asset's attributes using map
func patchAssetAttributes(a *Asset, attributes interface{}) {
	if attributes == nil {
		return
	}
	attributesMap, ok := attributes.(map[string]interface{})
	if !ok {
		return
	}

	if a.Attributes == nil {
		a.Attributes = attributesMap
		return
	}

	a.Attributes = mergemap.Merge(a.Attributes, attributesMap)
}

func patchString(key string, data map[string]interface{}, defaultVal string) string {
	_, exists := data[key]
	if !exists {
		return defaultVal
	}

	return getString(key, data)
}

func getString(key string, data map[string]interface{}) string {
	val, exists := data[key]
	if !exists {
		return ""
	}
	stringVal, ok := val.(string)
	if !ok {
		return ""
	}

	return stringVal
}
