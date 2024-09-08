CREATE MIGRATION m157xxf3vlgqxssnrf2bmswxhfrxucsvio6kyekk7jdoolctfs66ra
    ONTO initial
{
  CREATE TYPE default::Movie {
      CREATE REQUIRED PROPERTY created_at: cal::local_datetime {
          SET default := (cal::to_local_datetime(std::datetime_current(), 'UTC'));
      };
      CREATE REQUIRED PROPERTY genres: array<std::str>;
      CREATE REQUIRED PROPERTY runtime: std::int32;
      CREATE REQUIRED PROPERTY title: std::str;
      CREATE REQUIRED PROPERTY year: std::int32;
  };
};
