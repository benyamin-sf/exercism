// @ts-check

/**
 * The number of minutes it takes to prepare a single layer.
 */
const PREPARATION_MINUTES_PER_LAYER = 2;

/**
 * The expected oven time in minutes.
 */
export const EXPECTED_MINUTES_IN_OVEN = 40;

const PREP_TIME_FOR_LAYER = 2;

/**
 * Determines the number of minutes the lasagna still needs to remain in the
 * oven to be properly prepared.
 *
 * @param {number} actualMinutesInOven
 * @returns {number} the number of minutes remaining
 */
export function remainingMinutesInOven(actualMinutesInOven) {
  if (actualMinutesInOven > EXPECTED_MINUTES_IN_OVEN) {
    return 0;
  }

  return EXPECTED_MINUTES_IN_OVEN - actualMinutesInOven;
}

/**
 * Given a number of layers, determines the total preparation time.
 *
 * @param {number} numberOfLayers
 * @returns {number} the total preparation time
 */
export function preparationTimeInMinutes(numberOfLayers) {
  return numberOfLayers * PREP_TIME_FOR_LAYER;
}

/**
 * Calculates the total working time. That is, the time to prepare all the layers
 * of lasagna, and the time already spent in the oven.
 *
 * @param {number} numberOfLayers
 * @param {number} actualMinutesInOven
 * @returns {number} the total working time
 */
export function totalTimeInMinutes(numberOfLayers, actualMinutesInOven) {
  return (numberOfLayers * PREP_TIME_FOR_LAYER) + actualMinutesInOven;
}
