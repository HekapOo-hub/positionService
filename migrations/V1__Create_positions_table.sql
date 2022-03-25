create table positions (
                           id varchar(60),
                           account_id varchar(60),
                           order_id varchar(60),
                           open_price double precision,
                           close_price double precision,
                           take_profit double precision,
                           stop_loss double precision,
                           symbol varchar(10),
                           guaranteed_stop_loss bool,
                           state varchar(10),
                           quantity double precision,
                           leverage bool,
                           side varchar(10)
);


CREATE OR REPLACE FUNCTION notify_event() RETURNS TRIGGER AS $$

DECLARE
data json;
    notification json;

BEGIN

    data = row_to_json(NEW);

    -- Construct the notification as a JSON string.
    notification = json_build_object(
            'action', TG_OP,
            'position', data);


    -- Execute pg_notify(channel, notification)
    PERFORM pg_notify('events',notification::text);

    -- Result is ignored since this is an AFTER trigger
RETURN NULL;
END;

$$ LANGUAGE plpgsql;

CREATE TRIGGER position_notify_event
    AFTER INSERT OR UPDATE OR DELETE ON positions
    FOR EACH ROW EXECUTE PROCEDURE notify_event();