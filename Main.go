package main

import (
	"flag"
	"fmt"
	"go-blockchain/blockchain"
	"os"
	"runtime"
)

type CommandLine struct {
	blockchain *blockchain.BlockChain
}

func (cli *CommandLine) printUsage() {
	fmt.Println("Usage: ")

}

func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

func (cli *CommandLine) addBlock(data string) {
	cli.blockchain.AddBlock(data)
	fmt.Println("Added block")
}

func (cli *CommandLine) printChain() {
	iter := cli.blockchain.Iterator()
	for {
		block := iter.Next()
		fmt.Printf("PrevHash %x\n", block.PrevHash)
		fmt.Printf("Data %s\n", block.Data)
		fmt.Printf("Hash %x\n", block.Hash)
		fmt.Println("===========")
    if (len(block.PrevHash)==0){
      break
    }
	}
}

func (cli *CommandLine) run(){
  cli.validateArgs()
  addBlockCmd := flag.NewFlagSet("add",flag.ExitOnError)
  printChainCmd :=flag.NewFlagSet("print", flag.ExitOnError)
  addBlockData := addBlockCmd.String("block", "", "Block data")

  switch os.Args[1]{
    case "add":
      err:=addBlockCmd.Parse(os.Args[2:])
      blockchain.Handle(err)

      
      case "print":
        err:=printChainCmd.Parse(os.Args[2:])
        blockchain.Handle(err)
      default:
        cli.printUsage()
        runtime.Goexit()
    }
    if addBlockCmd.Parsed(){
      if *addBlockData == ""{
        addBlockCmd.Usage()
        runtime.Goexit()
      }
      cli.addBlock(*addBlockData)
    }
    if printChainCmd.Parsed(){
      cli.printChain()
    }

}

func main() {
	// chain := blockchain.InitBlockChain()
	// chain.AddBlock("First")
	// chain.AddBlock("Second")
	// chain.AddBlock("Third")

	// for _,block :=range chain.Blocks {
	// 	fmt.Printf("PrevHash %x\n", block.PrevHash)
	// 	fmt.Printf("Data %s\n", block.Data)
	// 	fmt.Printf("Hash %x\n",block.Hash)

	// 	fmt.Println("===========")
	// }
  defer os.Exit(0)
  chain :=blockchain.InitBlockChain()
  defer chain.Database.Close()
  cli:=CommandLine{chain}
  cli.run()

}
