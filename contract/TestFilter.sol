//SPDX-License-Identifier: Unlicense
pragma solidity ^0.4.23;

contract TestFilter {

    string name;

    uint256 age;

    event TestSetName(string oldName,string newName,uint256 timestamp);

    event TestSetAge(uint256 oldAge,uint256 newAge,uint256 timestamp);

    event TestGetName(string funcName,uint256 blockNumber);

    event TestGetAge(string funcName,uint256 blockNumber);

    constructor(string  memory name_,uint256 age_) public {
        name = name_;
        age = age_;
        emit TestSetName("", name, block.timestamp);
        emit TestSetAge(0,age,block.timestamp);
    }

    function setName(string _name) public {
        emit TestSetName(name, _name, block.timestamp);
        name = _name;  
    }

    
    function setAge(uint256 _age) public {
        emit TestSetAge(age, _age, block.timestamp);
        age = _age;  
    }

    function getName()public payable returns (string) {
        emit TestGetName("getName()", block.number);
        return name;
    }


    function getAge()public payable returns (uint256) {
         emit TestGetAge("getAge()", block.number);
        return age;
    }
}