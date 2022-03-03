let sheet = document.getElementById('plot');
let data = [];

Object.entries(plot).forEach((item) => {
    data.push({
        x : item[1].X,
        y : item[1].Y,
        type : 'scatter',
        name : item[0]
    })
})

const layout = {
    title : "Результаты эксперимента",
    xaxis: {
        title: 'Длина сообщения, Б',
        showgrid: true,
        zeroline: true
    },
    yaxis: {
        title: 'Продолжительность, мкс',
        showline: true,
        zeroline: true
    }
};

Plotly.newPlot(sheet, data, layout);