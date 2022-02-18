package handlers

import (
	"errors"
	"net/http"
	"net/mail"

	"github.com/gorilla/mux"
	"github.com/odpf/columbus/asset"
	"github.com/odpf/columbus/star"
	"github.com/odpf/columbus/user"
	"github.com/odpf/salt/log"
)

// UserHandler exposes a REST interface to user
type UserHandler struct {
	starRepository star.Repository
	logger         log.Logger
}

func (h *UserHandler) GetStarredAssetsWithHeader(w http.ResponseWriter, r *http.Request) {
	userID := user.FromContext(r.Context())
	if userID == "" {
		h.logger.Warn(errMissingUserID.Error())
		WriteJSONError(w, http.StatusBadRequest, errMissingUserID.Error())
		return
	}

	starCfg := buildStarConfig(h.logger, r.URL.Query())

	starredAssets, err := h.starRepository.GetAllAssetsByUserID(r.Context(), starCfg, userID)
	if err != nil {
		if errors.Is(err, star.ErrEmptyUserID) || errors.As(err, new(star.InvalidError)) {
			WriteJSONError(w, http.StatusBadRequest, err.Error())
			return
		}
		if errors.As(err, new(star.NotFoundError)) {
			WriteJSONError(w, http.StatusNotFound, err.Error())
			return
		}
		internalServerError(w, h.logger, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, starredAssets)
}

func (h *UserHandler) GetStarredAssetsWithPath(w http.ResponseWriter, r *http.Request) {
	targetUserID := mux.Vars(r)["user_id"]
	if targetUserID == "" {
		WriteJSONError(w, http.StatusBadRequest, errMissingUserID.Error())
		return
	}

	starCfg := buildStarConfig(h.logger, r.URL.Query())

	var starredAssets []asset.Asset

	// we can use user id or user email
	// a temporary flow and might be deleted in the future version
	_, err := mail.ParseAddress(targetUserID)
	if err == nil {
		starredAssets, err = h.starRepository.GetAllAssetsByUserEmail(r.Context(), starCfg, targetUserID)
	} else {
		starredAssets, err = h.starRepository.GetAllAssetsByUserID(r.Context(), starCfg, targetUserID)
	}
	if err != nil {
		if errors.Is(err, star.ErrEmptyUserID) || errors.As(err, new(star.InvalidError)) {
			WriteJSONError(w, http.StatusBadRequest, err.Error())
			return
		}
		if errors.As(err, new(star.NotFoundError)) {
			WriteJSONError(w, http.StatusNotFound, err.Error())
			return
		}
		internalServerError(w, h.logger, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, starredAssets)
}

func (h *UserHandler) StarAsset(w http.ResponseWriter, r *http.Request) {
	userID := user.FromContext(r.Context())
	if userID == "" {
		h.logger.Warn(errMissingUserID.Error())
		WriteJSONError(w, http.StatusBadRequest, errMissingUserID.Error())
		return
	}

	pathParams := mux.Vars(r)
	assetID := pathParams["asset_id"]

	starID, err := h.starRepository.Create(r.Context(), userID, assetID)
	if err != nil {
		if errors.Is(err, star.ErrEmptyAssetID) || errors.Is(err, star.ErrEmptyUserID) || errors.As(err, new(star.InvalidError)) {
			WriteJSONError(w, http.StatusBadRequest, err.Error())
			return
		}
		if errors.As(err, new(star.UserNotFoundError)) {
			WriteJSONError(w, http.StatusNotFound, err.Error())
			return
		}
		if errors.As(err, new(star.DuplicateRecordError)) {
			// idempotent
			writeJSON(w, http.StatusNoContent, starID)
			return
		}
		internalServerError(w, h.logger, err.Error())
		return
	}

	writeJSON(w, http.StatusNoContent, starID)
}

func (h *UserHandler) GetStarredAsset(w http.ResponseWriter, r *http.Request) {
	userID := user.FromContext(r.Context())
	if userID == "" {
		h.logger.Warn(errMissingUserID.Error())
		WriteJSONError(w, http.StatusBadRequest, errMissingUserID.Error())
		return
	}

	pathParams := mux.Vars(r)
	assetID := pathParams["asset_id"]

	starID, err := h.starRepository.GetAssetByUserID(r.Context(), userID, assetID)
	if err != nil {
		if errors.Is(err, star.ErrEmptyAssetID) || errors.Is(err, star.ErrEmptyUserID) || errors.As(err, new(star.InvalidError)) {
			WriteJSONError(w, http.StatusBadRequest, err.Error())
			return
		}
		if errors.As(err, new(star.NotFoundError)) {
			WriteJSONError(w, http.StatusNotFound, err.Error())
			return
		}
		internalServerError(w, h.logger, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, starID)
}

func (h *UserHandler) UnstarAsset(w http.ResponseWriter, r *http.Request) {
	userID := user.FromContext(r.Context())
	if userID == "" {
		h.logger.Warn(errMissingUserID.Error())
		WriteJSONError(w, http.StatusBadRequest, errMissingUserID.Error())
		return
	}

	pathParams := mux.Vars(r)
	assetID := pathParams["asset_id"]

	err := h.starRepository.Delete(r.Context(), userID, assetID)
	if err != nil {
		if errors.Is(err, star.ErrEmptyAssetID) || errors.Is(err, star.ErrEmptyUserID) || errors.As(err, new(star.InvalidError)) {
			WriteJSONError(w, http.StatusBadRequest, err.Error())
			return
		}
		if errors.As(err, new(star.NotFoundError)) {
			WriteJSONError(w, http.StatusNotFound, err.Error())
			return
		}
		internalServerError(w, h.logger, err.Error())
		return
	}

	writeJSON(w, http.StatusNoContent, "success")
}

func NewUserHandler(logger log.Logger, starRepo star.Repository) *UserHandler {
	h := &UserHandler{
		starRepository: starRepo,
		logger:         logger,
	}
	return h
}
