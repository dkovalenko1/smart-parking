--Up

CREATE TABLE IF NOT EXISTS parking_zones (
    id          UUID        PRIMARY KEY,
    name        TEXT        NOT NULL UNIQUE,
    description TEXT        NOT NULL DEFAULT 'Some parking zone',
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS spots (
     id         UUID        PRIMARY KEY,
     zone_id    UUID        NOT NULL REFERENCES parking_zones(id),
     number     TEXT        NOT NULL,
     type       TEXT        NOT NULL CHECK (type IN ('regular', 'ev', 'disabled')),
     status     TEXT        NOT NULL DEFAULT 'available' CHECK (status IN ('available', 'reserved', 'occupied', 'maintenance')),
     created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
     UNIQUE (zone_id, number)
);

CREATE TABLE IF NOT EXISTS reservations (
    id         UUID        PRIMARY KEY,
    spot_id    UUID        NOT NULL REFERENCES spots(id),
    user_id    UUID        NOT NULL,
    start_time TIMESTAMPTZ NOT NULL,
    end_time   TIMESTAMPTZ NOT NULL,
    status     TEXT        NOT NULL DEFAULT 'created' CHECK (status IN ('created', 'checked_in', 'checked_out', 'cancelled')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS bills (
    id               UUID          PRIMARY KEY,
    user_id          UUID          NOT NULL,
    reservation_id   UUID          NOT NULL REFERENCES reservations(id),
    duration_minutes INT           NOT NULL CHECK (duration_minutes > 0),
    total            NUMERIC(10,2) NOT NULL,
    currency         TEXT          NOT NULL DEFAULT 'UAH',
    status           TEXT          NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'paid', 'failed', 'cancelled')),
    created_at       TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    paid_at          TIMESTAMPTZ
);

