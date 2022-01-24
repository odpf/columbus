package handlers

import (
	"fmt"
	"github.com/odpf/salt/log"
	"net/http"
	"strings"

	"github.com/odpf/columbus/record"
)

// TypeHandler exposes a REST interface to types
type TypeHandler struct {
	typeRepo record.TypeRepository
	logger   log.Logger
}

func NewTypeHandler(logger log.Logger, er record.TypeRepository) *TypeHandler {
	h := &TypeHandler{
		typeRepo: er,
		logger:   logger,
	}

	return h
}

func (h *TypeHandler) Get(w http.ResponseWriter, r *http.Request) {
	typesNameMap, err := h.typeRepo.GetAll(r.Context())
	if err != nil {
		internalServerError(w, h.logger, "error fetching types")
		return
	}

	type TypeWithCount struct {
		Name  string `json:"name"`
		Count int    `json:"count"`
	}

	results := []TypeWithCount{}
	for _, typName := range record.AllSupportedTypes {
		count, _ := typesNameMap[typName]
		results = append(results, TypeWithCount{
			Name:  typName.String(),
			Count: count,
		})
	}

	writeJSON(w, http.StatusOK, results)
}

func (h *TypeHandler) parseSelectQuery(raw string) (fields []string) {
	tokens := strings.Split(raw, ",")
	for _, token := range tokens {
		token = strings.TrimSpace(token)
		if token == "" {
			continue
		}
		fields = append(fields, token)
	}
	return
}

func (h *TypeHandler) selectRecordFields(fields []string, records []record.Record) (processedRecords []record.Record) {
	for _, record := range records {
		newData := map[string]interface{}{}
		for _, field := range fields {
			v, ok := record.Data[field]
			if !ok {
				continue
			}
			newData[field] = v
		}
		record.Data = newData
		processedRecords = append(processedRecords, record)
	}
	return
}

func (h *TypeHandler) validateRecord(record record.Record) error {
	if record.Urn == "" {
		return fmt.Errorf("urn is required")
	}
	if record.Name == "" {
		return fmt.Errorf("name is required")
	}
	if record.Data == nil {
		return fmt.Errorf("data is required")
	}
	if record.Service == "" {
		return fmt.Errorf("service is required")
	}

	return nil
}

func (h *TypeHandler) responseStatusForError(err error) (int, string) {
	switch err.(type) {
	case record.ErrNoSuchType, record.ErrNoSuchRecord:
		return http.StatusNotFound, err.Error()
	}
	return http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)
}
