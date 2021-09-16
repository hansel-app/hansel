// TODO: replace with interface generated from protobuf?
export interface Gem {
  color: GemColor;
}

export interface GemColor {
  selectionColor: HexCode;
  tintColor: HexCode;
}

type HexCode = String;
