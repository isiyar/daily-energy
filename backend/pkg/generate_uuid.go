package utils

import (
  "github.com/google/uuid"
  "github.com/jackc/pgx/v5/pgtype"
)

func GenerateUUID() pgtype.UUID {
  u := uuid.New()
  return pgtype.UUID{
    Bytes: u,
    Valid: true,
  }
}
