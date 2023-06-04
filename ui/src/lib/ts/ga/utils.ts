import * as _ from "lodash";

export class Utils {
  static generateRandomChars(length: number): string[] {
    const characterSet: string[] = ["0", "1"];
    const randomChars: string[] = _.times(length, () => _.sample(characterSet) || "");;
    return randomChars;
  }
}