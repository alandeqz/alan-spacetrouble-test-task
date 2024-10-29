-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS "alan-tabeo-test-task"."bookings"
(
    id             BIGSERIAL PRIMARY KEY,
    created_at     TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    first_name     VARCHAR(255)                NOT NULL,
    last_name      VARCHAR(255)                NOT NULL,
    gender         SMALLINT                    NOT NULL,
    birthday       DATE                        NOT NULL,
    launchpad_id   VARCHAR(255)                NOT NULL,
    destination_id VARCHAR(255)                NOT NULL,
    launch_date    TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

COMMENT ON COLUMN "alan-tabeo-test-task"."bookings".gender IS 'Unknown=0;Male=1;Female=2;Other=3';

CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_booking_updated_at
    BEFORE UPDATE
    ON "alan-tabeo-test-task"."bookings"
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin

DROP TRIGGER IF EXISTS update_booking_updated_at ON "alan-tabeo-test-task"."bookings";

DROP FUNCTION IF EXISTS update_updated_at_column;

DROP TABLE IF EXISTS "alan-tabeo-test-task"."bookings";

-- +goose StatementEnd
