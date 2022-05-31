//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

contract TestFilter {

    string  name;

    uint256 age;

    event TestSetName(string indexed oldName, string indexed newName,uint256 indexed timestamp);

    event TestSetAge(uint256 indexed oldAge,uint256 indexed newAge,uint256 indexed timestamp);

    event TestGetName(string funcName,uint256 blockNumber);

    event TestGetAge(string funcName,uint256 blockNumber);

    constructor(string  memory name_,uint256 age_)  {
        name = name_;
        age = age_;
        emit TestSetName("", name, block.timestamp);
        emit TestSetAge(0,age,block.timestamp);
    }

    function setName(string memory _name) public {
        emit TestSetName(name, _name, block.timestamp);
        name = _name;  
    }

    
    function setAge(uint256 _age) public {
        emit TestSetAge(age, _age, block.timestamp);
        age = _age;  
    }

    function getName()public {
        emit TestGetName("getName()", block.number);
    }


    function getAge()public payable returns (uint256) {
         emit TestGetAge("getAge()", block.number);
        return age;
    }
}