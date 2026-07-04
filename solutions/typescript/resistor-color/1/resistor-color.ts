export const colorCode = (bandColor: string) => {
  return COLORS.findIndex(color => color === bandColor)
}

export const COLORS = ["black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"]
