export type Function = {
  title: string,
  description: string,
  name: string,
  placeholder: string,
  inputs: Variable[],
  outputs: Variable[]
}

type Variable = {
  name: string,
  type: string,
}