function hello(name: string) :string {
  console.log(`===>hello ${name || "bro"}`);
  return `hello ${name || 'bro'}`
}

hello("BBB");
hello("HHH");
hello("YYY");