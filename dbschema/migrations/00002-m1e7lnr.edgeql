CREATE MIGRATION m1e7lnrvlz7hfwzpifktb6sqbmxy2yjg2k3ptx4xxy5osrunebmnwq
    ONTO m157xxf3vlgqxssnrf2bmswxhfrxucsvio6kyekk7jdoolctfs66ra
{
  ALTER TYPE default::Movie {
      CREATE REQUIRED PROPERTY version: std::int32 {
          SET default := 1;
      };
  };
};
