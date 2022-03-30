function getTableValuesById(table_id){
    table = document.getElementById(table_id)
    let output = [];
    for(let i = 1; i < table.rows.length; i++) {
        tmp = {"ID" : parseInt(table.rows[i].cells[2].innerHTML), "Name" : table.rows[i].cells[0].innerHTML};
        output.push(tmp);
    }
    return output;
}

function getInputDataAttrFromElemId(elem_id){
    elem = document.getElementById(elem_id)
    return [elem.value, parseInt(elem.dataset.id)]
}

function getDropDownDataAttrFromElemId(elem_id){
    elem = document.getElementById(elem_id)
    return [elem.value, parseInt(elem.options[elem.selectedIndex].dataset.id)]
}

function showModal(modal_id, title, content){
    var myModal = new bootstrap.Modal(document.getElementById(modal_id))
    document.getElementById('modalTitle').innerHTML = title
    document.getElementById('modalContent').innerHTML = content
    myModal.show()      
}

function changeElementClass(element_id, old_class, new_class) {
    elem = document.getElementById(element_id)
    elem.classList.remove(old_class);
    elem.classList.add(new_class)
}

function onAddJob() {
    programming_skills = getTableValuesById("programming-skill-table")
    if (programming_skills.length === 0) {
        showModal("add_job_modal", "But please...", "Include a programming skill... it wouldn't be a legal job position or something like that, right? hehe")
        return
    }
    title = document.getElementById("title").value
    description = document.getElementById("description").value
    link = document.getElementById("link").value
    publishedAt = new Date(document.getElementById("publishedAt").value).toISOString()
    finished = document.getElementById("finished").value
    let [company_name, company_id] = getInputDataAttrFromElemId("companies")
    let [publisher_name, publisher_id] = getInputDataAttrFromElemId("publishers")
    let [company_size, company_size_id] = getDropDownDataAttrFromElemId("company-sizes")
    let [english_lvl, english_lvl_id] = getDropDownDataAttrFromElemId("english-levels")
    let [work_schedule, work_schedule_id] = getDropDownDataAttrFromElemId("schedules")
    years_of_xp = parseInt(document.getElementById("experience").value)
    salary = parseInt(document.getElementById("salary").value)
    personal_skills = getTableValuesById("personal-skill-table")

    output = {
        Title : title,
        Description: description,
        PublisherReferer : publisher_id,
        Publisher: {
            ID: publisher_id,
            Name: publisher_name
        },
        Link: link,
        CompanyReferer: company_id,
        Company: {
            ID: company_id,
            Name: company_name,
            SizeReferer: company_size_id,
            Size: {
                ID: company_size_id,
                Name: company_size
            }
        },
        PublishedAt: publishedAt,
        EnglishLevelReferer : english_lvl_id,
        EnglishLevel: {
            ID: english_lvl_id,
            Level: english_lvl
        },
        Experience: years_of_xp,
        SchedulesReferer: work_schedule_id,
        Schedules: {
            ID: work_schedule_id,
            Name: work_schedule
        },
        ProgrammingSkills: programming_skills,
        PersonalSkills: personal_skills,
        Salary: salary
    };

    //// Dummy data
    // output = {
    //     Title : "title",
    //     Description: "description",
    //     PublisherReferer : 2,
    //     Publisher: {
    //         ID: 2,
    //         Name: "InfoJobs"
    //     },
    //     Link: "link",
    //     CompanyReferer: 1,
    //     Company: {
    //         ID: 1,
    //         Name: "CloudBlue",
    //         SizeReferer: 1,
    //         Size: {
    //             ID: 1,
    //             Name: "Micro"
    //         }
    //     },
    //     PublishedAt: "2022-03-02T00:00:00.000Z",
    //     EnglishLevelReferer : 6,
    //     EnglishLevel: {
    //         ID: 6,
    //         Level: "Beginner"
    //     },
    //     Experience: 6,
    //     SchedulesReferer: 1,
    //     Schedules: {
    //         ID: 1,
    //         Name: "Part-time"
    //     },
    //     ProgrammingSkills: [
    //         {
    //             ID: 11,
    //             Name: "JAVADOC"
    //         },
    //         {
    //             ID: 12,
    //             Name: "NodeJS"
    //         }],
    //     PersonalSkills: [
    //         {
    //             ID: 1,
    //             Name: "Responsibility"
    //         },
    //         {
    //             ID: 2,
    //             Name: "Attention to detail"
    //         }],
    //     Salary: 9000
    // };

    if (!document.getElementById("finished").checked) {
        output.FinishedAt = new Date(document.getElementById("finishedAt").value).toISOString()
    }
    
    //Send POST request to server
    fetch("http://localhost:3000/addjob", {
        method: "POST",
        body: JSON.stringify(output),
        headers: {
            'Content-Type': 'application/json'
          }
        })
        .then(response => response.json())
        .then(result => {
            if (result.Status === 200){
                changeElementClass("close_modal_bttn", "btn-secondary", "btn-primary")
                showModal('add_job_modal', 'Congratulations', 'Job titled "' + title + '" was successfully inserted!')   
            } else{
                showModal('add_job_modal', 'Error Status: ' + result.Status, 'There was an error: ' + result.Msg)      
            }
        })
}

var tables = document.getElementsByTagName("table")

for(i = 0; i < document.getElementsByName("tags").length; i++){
    tables[i].addEventListener("click", onDeleteRow);
}

function addRow(table, id, value){
    // find the way ive got the element, now lets find a subtag 
    tbody = table.getElementsByTagName("tbody")[0];
    tbody.innerHTML += `
        <tr>
            <td>${value}</td>
            <td><button type="button" class="btn btn-danger btn-sm">Delete</button></td>
            <td>${id}</td>
        </tr>
    `;
}

function onAddNewProgSkillBtnClick(input_id, table_id) {
    value = document.getElementById(input_id).value
    if(value) addRow(document.getElementById(table_id), 0, value)
}
    
function onDeleteRow(e) {
    if (!e.target.classList.contains("btn")) {
        return;
    }
    const btn = e.target;
    btn.closest("tr").remove();
}

function getAllData(input_id, table_id = null) {
    let field = document.getElementById(input_id);
    let ac = new Autocomplete(field, {
        maximumItems: 50,
        threshold: 1,
        onSelectItem: ({label, value}) => {
            console.log("user selected:", label, value);
            if (table_id){
                addRow(document.getElementById(table_id), value, label)
            } else {
                field.setAttribute('data-id', value);
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
                    tmp = {"label" : data[i].Name, "value" : data[i].ID};
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
getAllData("companies")
getAllData("publishers")
getAllData("programming-skill-names", "programming-skill-table")
getAllData("personal-skill-names", "personal-skill-table")

//make datapicker only show past days
function preventFutureDatesDatapicker(datapicker_id){
    var dtToday = new Date();
    var month = dtToday.getMonth() + 1;
    var day = dtToday.getDate();
    var year = dtToday.getFullYear();
    if(month < 10)
        month = '0' + month.toString();
    if(day < 10)
        day = '0' + day.toString();
    var maxDate = year + '-' + month + '-' + day;
    document.getElementById(datapicker_id).setAttribute("max", maxDate);
}

// Show finished datapicker based on checkbox result
function finishedValueChanged(display_div_id)
{   
    var visualization = document.getElementById(display_div_id).style.display;
    if(visualization === "block"){
        visualization = "none"
        document.getElementById("finishedAt").required = false;
    } else {
        visualization = "block"
        document.getElementById("finishedAt").required = true;
    }
    document.getElementById(display_div_id).style.display = visualization;
}

preventFutureDatesDatapicker("publishedAt")
preventFutureDatesDatapicker("finishedAt")

const element = document.querySelector('form');
element.addEventListener('submit', event => {
  event.preventDefault();
  // actual logic, e.g. validate the form
  console.log('Form submission cancelled.');
});

