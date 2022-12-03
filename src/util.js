export function shuffleArray(array) {
  for (let i = array.length - 1; i > 0; i--) {
    const j = Math.floor(Math.random() * (i + 1));
    [array[i], array[j]] = [array[j], array[i]];
  }
}

export function factorial(n) {
  let ans = 1;
  for (let i = 2; i <= n; i++) {
    ans = ans * i;
  }
  return ans;
}

// Determine if any gifters are giving to someone who
// is also giving back to them.
export function circularPairing(giftersToRecievers) {
  for (const firstPerson in giftersToRecievers) {
    const secondPerson = giftersToRecievers[firstPerson];
    const thirdPerson = giftersToRecievers[secondPerson];
    if (firstPerson === thirdPerson) {
      return true;
    }
  }
  return false;
}
