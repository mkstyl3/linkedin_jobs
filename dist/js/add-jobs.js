function getTableValuesById(table_id){
    table = document.getElementById(table_id)
    values = []
    for(i=1; i<table.rows.length; i++){
        values.push(table.rows[i].cells[0].innerHTML)
    }
    return values
}

function onAddJob() {
    programming_skills = getTableValuesById("programming-skill-table")
    if (programming_skills.length === 0) {
        alert("Include a programming skill... it wouldn't be a legal job position or something like that, right? hehe");
        return
    }
    title = document.getElementById("title").value
    publishers = document.getElementById("publishers").value
    description = document.getElementById("description").value
    link = document.getElementById("link").value
    publishedAt = document.getElementById("publishedAt").value
    finishedAt = document.getElementById("finishedAt").value
    company_name = document.getElementById("company-names").value
    company_size = document.getElementById("company-sizes").value
    english_lvl = document.getElementById("english-levels").value
    work_schedules = document.getElementById("schedules").value
    salary = document.getElementById("salary").value
    years_of_xp = document.getElementById("experience").value
    personal_skills = getTableValuesById("personal-skill-table")
    // output = {
    //     "Title" : title,
    //     "Description": description,
    //     "Publisher" : publishers,
    //     "Link": link,
    //     "PublishedAt": publishedAt,
    //     "FinishedAt": finishedAt,
    //     "Company" : {
    //         "Name" : "",
    //         "Size" : {
                
    //         }
    //     }
    //     "company_size" : company_size,
    //     "english_lvl" : english_lvl,
    //     "work_schedules": work_schedules,
    //     "salary": salary,
    //     "years_of_xp": years_of_xp,
    //     "programming_skills": JSON.stringify(programming_skills),
    //     "personal_skills": JSON.stringify(personal_skills),
    // };
    // Send POST request to server
    // fetch("http://localhost:3000/addjob", {method: "POST", body: output})
    //     .then(results => results.json())
    //     .then(console.log);

}

var tables = document.getElementsByTagName("table")

for(i = 0; i < document.getElementsByName("tags").length; i++){
    tables[i].addEventListener("click", onDeleteRow);
}

function addRow(table, value){
    // find the way ive got the element, now lets find a subtag 
    tbody = table.getElementsByTagName("tbody")[0];
    tbody.innerHTML += `
        <tr>
            <td>${value}</td>
            <td><button type="button" class="btn btn-danger btn-sm">Delete</button></td>
        </tr>
    `;
}

function onAddNewProgSkillBtnClick(input_id, table_id) {
    input = document.getElementById(input_id).value
    if(input) addRow(document.getElementById(table_id), input)
}
    
function onDeleteRow(e) {
    if (!e.target.classList.contains("btn")) {
        return;
    }
    const btn = e.target;
    btn.closest("tr").remove();
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
let table = document.getElementById('programming-skill-table');
let ac2 = new Autocomplete(field2, {
    maximumItems: 50,
    threshold: 1,
    onSelectItem: ({label, value}) => {
        console.log("user selected:", label, value);
        addRow(table, field2.value)
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
let table2 = document.getElementById('personal-skill-table');
let ac3 = new Autocomplete(field3, {
    maximumItems: 50,
    threshold: 1,
    onSelectItem: ({label, value}) => {
        console.log("user selected:", label, value);
        addRow(table2, field3.value)
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

function getAllData(input_id, table_id = null) {
    let field = document.getElementById(input_id);
    let ac = new Autocomplete(field, {
        maximumItems: 50,
        threshold: 1,
        onSelectItem: ({label, value}) => {
            console.log("user selected:", label, value);
            if (table_id){
                addRow(document.getElementById(table_id), field.value)
            }
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
}

getAllData("publishers")

//make datapicker only show past days
var dtToday = new Date();
var month = dtToday.getMonth() + 1;
var day = dtToday.getDate();
var year = dtToday.getFullYear();
if(month < 10)
    month = '0' + month.toString();
if(day < 10)
    day = '0' + day.toString();
var maxDate = year + '-' + month + '-' + day;
document.getElementById("publishedAt").setAttribute("max", maxDate);