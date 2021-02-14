import {React, useRef, useEffect, useState} from 'react';
import * as d3 from "d3";
import "../styles/Graph.css";

const ToolTip = (props) => {
    if (props.left === 0 || props.top === 0) {
        return (null);
    }
    else {
        return (
            <div className="tooltip" style={{ left: `${props.left + 15}px`, top: `${props.top + 15}px` }}>
                {props.fields.map((field, i) => <p key={i}>{field}</p>)}
            </div>
        );
    }
}

const Graph = (props) => {
    const graphRef = useRef(null);
    const [topState, setTopState] = useState(0);
    const [leftState, setLeftState] = useState(0);
    const [fieldsState, setFieldsState] = useState([]);

    const maxVal = data => {
        let maxX = 0;
        let maxY = 0;
        console.log(data);
        data.forEach(e => {
            maxX = Math.max(maxX, Math.abs(e["x"]));
            maxY = Math.max(maxY, Math.abs(e["y"]));
        });
        
        return [maxX, maxY];
    }

    useEffect(() => {
            const width = 700;
            const height = 700;
            const margin = 50;
            let svg = d3.select(graphRef.current)
                .append("svg")
                .attr("width", width)
                .attr("height", height)
            const data = props.data;
            const [maxX, maxY] = maxVal(data);

            let x = d3.scaleLinear().domain([-maxX,maxX]).range([margin, width - margin]);
            let y = d3.scaleLinear().domain([-maxY,maxY]).range([height - margin, margin]);
    
            svg.selectAll("circle")
                .data(data).enter().append("circle")
                .style("fill", "orange")
                .attr("r", 10)
                .attr("cx", d => x(d.x))
                .attr("cy", d => y(d.y))
                .on("mouseover", d => {
                    setLeftState(d.x);
                    setTopState(d.y);
                    setFieldsState([
                        `Positivity: ${x(d.x)}`,
                        `Energy: ${x(d.y)}`,
                        `Title: ${d.title}`
                    ]);
                })
                .on("mouseout", d => {
                    setLeftState(0);
                    setTopState(0);
                    setFieldsState([]);
                });
            
            const xAxis = d3.axisBottom().tickValues([]).scale(x);
            const yAxis = d3.axisLeft().tickValues([]).scale(y);
            const xTranslate = height/2;
            const yTranslate = width/2;
            svg.append("g").attr("transform", "translate(0," +  xTranslate + ")").call(xAxis);
            svg.append("g").attr("transform", "translate(" + yTranslate + ", 0)").call(yAxis);
            d3.json("data.json")
            
            const xTextOffset = (height / 2) - 20;
            const yTextOffset = (width / 2) + 50;
            svg.append("text").text("Positivity")
                .attr("x", 30)
                .attr("y", xTextOffset)
                .style("text-anchor", "middle");
            svg.append("text").text("Energy")
                .attr("x", yTextOffset)
                .attr("y", 20)
                .style("text-anchor", "middle");

    }, [props.data]);

    return (
        <div>
            <div ref={graphRef} className="main"/>
            <ToolTip left={leftState} top={topState} fields={fieldsState}/>
        </div>
    );
}

export default Graph;