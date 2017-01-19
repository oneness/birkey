actor Main
  new create(env: Env) =>
    env.out.print("Hello, world!")
    var a = Aardvark.create("eater")
    var b = a.eat(4)
    //env.out.print(b)
    
actor Aardvark
  let name: String
  var _hunger_level: U64 = 0

  new create(name': String) =>
    name = name'

  be eat(amount: U64) =>
    _hunger_level = _hunger_level - amount.min(_hunger_level)    
    
