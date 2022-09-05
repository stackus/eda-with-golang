-- Apply to keep modifications to the created_at column from being made
CREATE OR REPLACE FUNCTION created_at_trigger() RETURNS TRIGGER AS
$$
BEGIN
  NEW.created_at := OLD.created_at;
  RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';

-- Apply to a table to automatically update update_at columns
CREATE OR REPLACE FUNCTION updated_at_trigger() RETURNS TRIGGER AS
$$
BEGIN
  IF ROW (NEW.*) IS DISTINCT FROM ROW (OLD.*) THEN
    NEW.updated_at = NOW();
    RETURN NEW;
  ELSE
    RETURN OLD;
  END IF;
END;
$$ LANGUAGE 'plpgsql';
