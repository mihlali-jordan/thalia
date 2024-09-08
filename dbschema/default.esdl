module default {
  type Movie {
      required title: str;
      required year: int32;
      required runtime: int32;
      required genres: array<str>;
      required version: int32 {
        default := 1;
      };
      required created_at: cal::local_datetime {
         default := cal::to_local_datetime(datetime_current(), 'UTC')
      };
  }
}
