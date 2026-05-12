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

INSERT INTO parking_zones (id, name, description)
VALUES
    ('11111111-1111-1111-1111-111111111111', 'Zone A', 'Main outdoor parking area'),
    ('22222222-2222-2222-2222-222222222222', 'Zone B', 'Secondary outdoor parking area'),
    ('33333333-3333-3333-3333-333333333333', 'Underground', 'Underground parking area')
ON CONFLICT (name) DO NOTHING;

INSERT INTO spots (id, zone_id, number, type, status)
VALUES
    ('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaa1', '11111111-1111-1111-1111-111111111111', 'A-01', 'regular', 'available'),
    ('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaa2', '11111111-1111-1111-1111-111111111111', 'A-02', 'regular', 'available'),
    ('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaa3', '11111111-1111-1111-1111-111111111111', 'A-03', 'ev', 'available'),
    ('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaa4', '11111111-1111-1111-1111-111111111111', 'A-04', 'disabled', 'available'),
    ('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbb1', '22222222-2222-2222-2222-222222222222', 'B-01', 'regular', 'available'),
    ('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbb2', '22222222-2222-2222-2222-222222222222', 'B-02', 'regular', 'reserved'),
    ('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbb3', '22222222-2222-2222-2222-222222222222', 'B-03', 'ev', 'available'),
    ('cccccccc-cccc-cccc-cccc-ccccccccccc1', '33333333-3333-3333-3333-333333333333', 'U-01', 'regular', 'available'),
    ('cccccccc-cccc-cccc-cccc-ccccccccccc2', '33333333-3333-3333-3333-333333333333', 'U-02', 'ev', 'occupied'),
    ('cccccccc-cccc-cccc-cccc-ccccccccccc3', '33333333-3333-3333-3333-333333333333', 'U-03', 'disabled', 'maintenance')
ON CONFLICT (zone_id, number) DO NOTHING;
