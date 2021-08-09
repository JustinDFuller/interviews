import { entity } from "@justindfuller/entity";

const defaults = entity.defaults({
  Name: entity.String(),
  EmailDomain: entity.String(),
});

export function New(input) {
  const data = defaults(input);

  return entity.New(data, New);
}
