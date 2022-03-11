const tbodyEl = document.querySelector("tbody");
const tableEl = document.querySelector("table");
// var tables = document.getElementsByTagName("table")
// var tbodies = document.getElementsByTagName("tbody")

// for(i = 0; i < document.getElementsByName("tags"); i++){
    
// }


function addRow(data){
    tbodyEl.innerHTML += `
        <tr>
            <td>${data}</td>
            <td><button type="button" class="btn btn-danger btn-sm">Delete</button></td>
        </tr>
    `;
}

function onDeleteRow(e) {
    if (!e.target.classList.contains("btn")) {
        return;
    }
    const btn = e.target;
    btn.closest("tr").remove();
}
tableEl.addEventListener("click", onDeleteRow);

function onAddNewProgSkillBtnClick(id) {
    console.log(id)
    const input = document.getElementById(id).value
    addRow(input)
}

// Retrieving company names
let field = document.getElementById('company-names');
let ac = new Autocomplete(field, {
    maximumItems: 50,
    threshold: 1,
    onSelectItem: ({label, value}) => {
        console.log("user selected:", label, value);
    }
});

let api = "http://localhost:3000/" + field.id

new Promise((resolve) => {
    fetch(api)
        .then((response) => response.json())
        .then((data) => {
            let output = [];
            for(let i = 0; i < data.length; i++) {
                tmp = {"label" : data[i].Name, "value" : i};
                output.push(tmp);
            }
            ac.setData(output)
            resolve(data);
        })
        .catch((error) => {
            console.error(error);
        });
    })

// Retrieving programming skill names
let field2 = document.getElementById('programming-skill-names');
let ac2 = new Autocomplete(field2, {
    maximumItems: 50,
    threshold: 1,
    onSelectItem: ({label, value}) => {
        console.log("user selected:", label, value);
        onAddProgrammingSkill(label)
    }
});

let api2 = "http://localhost:3000/" + field2.id

new Promise((resolve) => {
    fetch(api2)
        .then((response) => response.json())
        .then((data) => {
            let output = [];
            for(let i = 0; i < data.length; i++) {
                tmp = {"label" : data[i].Name, "value" : i};
                output.push(tmp);
            }
            ac2.setData(output)
            resolve(data);
        })
        .catch((error) => {
            console.error(error);
        });
})

// Retrieving programming skill names
let field3 = document.getElementById('personal-skill-names');
let ac3 = new Autocomplete(field3, {
    maximumItems: 50,
    threshold: 1,
    onSelectItem: ({label, value}) => {
        console.log("user selected:", label, value);
        onAddProgrammingSkill(label)
    }
});

let api3 = "http://localhost:3000/" + field3.id

new Promise((resolve) => {
    fetch(api3)
        .then((response) => response.json())
        .then((data) => {
            let output = [];
            for(let i = 0; i < data.length; i++) {
                tmp = {"label" : data[i].Name, "value" : i};
                output.push(tmp);
            }
            ac3.setData(output)
            resolve(data);
        })
        .catch((error) => {
            console.error(error);
        });
})
    

// const tableEl = document.querySelector("table");

// function onAddProgrammingSkill(data){
//     const tbodyEl = document.querySelector("tbody");
//     tbodyEl.innerHTML += `
//         <tr>
//             <td>${data}</td>
//             <td><button type="button" class="btn btn-danger btn-sm">Delete</button></td>
//         </tr>
//     `;
// }

// function onDeleteRow(e) {
//     console.log("HOLAAA")
//     if (!e.target.classList.contains("btn")) {
//         return;
//     }
//     const btn = e.target;
//     btn.closest("tr").remove();
// }

// function onAddNewProgSkillBtnClick() {
//     const input = document.getElementById('programming-skill-names').value
//     onAddProgrammingSkill(input)
// }

// tableEl.addEventListener("click", onDeleteRow);


